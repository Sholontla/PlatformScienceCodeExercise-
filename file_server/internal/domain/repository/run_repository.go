package repository

import "file_server/internal/domain/entity"

type RunRepository interface {
	RunProcess(interval int, frequency string, driver_demo_path string) entity.Message
}
