package toolserror

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func Wrap(area string, err error) error {
	return fmt.Errorf("[SolutionLabs] %s %w", area, err)
}

func Warning(area string, err error) {
	log.Warn(fmt.Errorf("[SolutionLabs] %s %w", area, err))
}
