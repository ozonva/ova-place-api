package repo_test

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/pressly/goose/v3"

	"github.com/ozonva/ova-place-api/internal/models"
	"github.com/ozonva/ova-place-api/internal/repo"
)

const (
	dbDriver = "postgres"
	dbString = "user=postgres dbname=places_api sslmode=disable password=postgres"
)

var _ = Describe("Repo", func() {
	var (
		err                 error
		migrationConnection *sql.DB
		repoConnection      *sqlx.DB
		dockerCli           *client.Client
		containerId         string
	)

	BeforeSuite(func() {
		dockerCli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Fatalf("cannot init docker client: %v", err)
		}

		ctx := context.Background()
		resp, err := dockerCli.ContainerCreate(ctx, &container.Config{
			Env:          []string{"POSTGRES_PASSWORD=postgres", "POSTGRES_DB=places_api"},
			Image:        "postgres",
			ExposedPorts: nat.PortSet{"5432": struct{}{}},
		}, &container.HostConfig{
			PortBindings: map[nat.Port][]nat.PortBinding{"5432": {{HostIP: "127.0.0.1", HostPort: "5432"}}},
		}, nil, nil, "pg-test")
		if err != nil {
			log.Fatalf("cannot create docker container: %v", err)
		}

		if err := dockerCli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
			log.Fatalf("cannot start docker container: %v", err)
		}

		time.Sleep(time.Second * 4)

		repoConnection, err = sqlx.Connect(dbDriver, dbString)
		if err != nil {
			log.Fatalf("cannot connect to db for repo: %v", err)
		}

		migrationConnection, err = sql.Open(dbDriver, dbString)
		if err != nil {
			log.Fatalf("cannot connect to db for migrations: %v", err)
		}

		err = migrationConnection.Ping()
		if err != nil {
			log.Fatalf("cannot ping to db for migrations: %v", err)
		}

		if err := goose.Up(migrationConnection, "../../db/migrations"); err != nil {
			log.Fatalf("cannot migrate: %v", err)
		}

		containerId = resp.ID

	})

	AfterSuite(func() {
		err := dockerCli.ContainerStop(context.TODO(), containerId, nil)
		if err != nil {
			log.Fatalf("cannot stop docker container: %v", err)
		}

		err = dockerCli.ContainerRemove(context.TODO(), containerId, types.ContainerRemoveOptions{})
		if err != nil {
			log.Fatalf("cannot remove docker container: %v", err)
		}
	})

	BeforeEach(func() {
		entities := []models.Place{
			{
				UserID: 1,
				Seat:   "seat1",
				Memo:   "memo1",
			},
		}
		_, err := repoConnection.NamedExec(`INSERT INTO places (user_id, memo, seat)
        VALUES (:user_id, :memo, :seat)`, entities)
		if err != nil {
			log.Fatalf("cannot seed: %v", err)
		}
	})

	AfterEach(func() {
		tables := []string{"places"}
		_, _ = repoConnection.Exec("SET FOREIGN_KEY_CHECKS=0;")

		for _, v := range tables {
			_, _ = repoConnection.Exec(fmt.Sprintf("TRUNCATE TABLE %s;", v))
		}

		_, _ = repoConnection.Exec("SET FOREIGN_KEY_CHECKS=1;")
	})

	Describe("Get places count", func() {
		Context("all is ok", func() {
			It("should return places count", func() {
				repoInstance := repo.NewRepo(repoConnection)

				value, err := repoInstance.TotalCount(context.TODO())

				gomega.Expect(value).To(gomega.Equal(uint64(1)))
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	Describe("Create place", func() {
		Context("all is ok", func() {
			It("should return created place id", func() {
				repoInstance := repo.NewRepo(repoConnection)

				value, err := repoInstance.AddEntity(context.TODO(), models.Place{
					Memo:   "test3",
					Seat:   "test3",
					UserID: 3,
				})

				gomega.Expect(value).To(gomega.Not(gomega.BeNil()))
				gomega.Expect(err).To(gomega.BeNil())

				gomega.Expect(getCount(repoConnection)).To(gomega.Equal(uint64(2)))
			})
		})
	})

	Describe("Multi create place", func() {
		Context("all is ok", func() {
			It("should not return error", func() {
				repoInstance := repo.NewRepo(repoConnection)

				err := repoInstance.AddEntities(context.TODO(), []models.Place{{
					Memo:   "test3",
					Seat:   "test3",
					UserID: 3,
				}})

				gomega.Expect(err).To(gomega.BeNil())

				gomega.Expect(getCount(repoConnection)).To(gomega.Equal(uint64(2)))
			})
		})
	})

	Describe("List place", func() {
		Context("all is ok", func() {
			It("should return entities", func() {
				repoInstance := repo.NewRepo(repoConnection)

				result, err := repoInstance.ListEntities(context.TODO(), 1, 0)

				gomega.Expect(result[0].UserID).To(gomega.Equal(uint64(1)))
				gomega.Expect(result[0].Memo).To(gomega.Equal("memo1"))
				gomega.Expect(result[0].Seat).To(gomega.Equal("seat1"))
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	Describe("Describe place", func() {
		Context("all is ok", func() {
			It("should return entity", func() {
				repoInstance := repo.NewRepo(repoConnection)

				result, err := repoInstance.DescribeEntity(context.TODO(), getFirst(repoConnection))

				gomega.Expect(result.UserID).To(gomega.Equal(uint64(1)))
				gomega.Expect(result.Memo).To(gomega.Equal("memo1"))
				gomega.Expect(result.Seat).To(gomega.Equal("seat1"))
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	Describe("Update place", func() {
		Context("all is ok", func() {
			It("should not return error", func() {
				repoInstance := repo.NewRepo(repoConnection)

				err := repoInstance.UpdateEntity(context.TODO(), getFirst(repoConnection), models.Place{
					Memo:   "test3",
					Seat:   "test3",
					UserID: 3,
				})

				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	Describe("Delete place", func() {
		Context("all is ok", func() {
			It("should not return error", func() {
				repoInstance := repo.NewRepo(repoConnection)

				err := repoInstance.RemoveEntity(context.TODO(), getFirst(repoConnection))

				gomega.Expect(getCount(repoConnection)).To(gomega.Equal(uint64(0)))
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

})

func getCount(repoConnection *sqlx.DB) uint64 {
	var count uint64
	err := repoConnection.Get(&count, "SELECT count(1) FROM places")
	if err != nil {
		log.Fatalf("cannot get count: %v", err)
	}
	return count
}

func getFirst(repoConnection *sqlx.DB) uint64 {
	var id uint64
	err := repoConnection.Get(&id, "SELECT id FROM places limit 1")
	if err != nil {
		log.Fatalf("cannot get first id: %v", err)
	}
	return id
}
