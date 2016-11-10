package chat

import "github.com/Sirupsen/logrus"

// Ctx ...
type Ctx struct {
	*Config
	*Logger
}

// NewCtx ... ConfigとLoggerの生成
func NewCtx(configpath string) (*Ctx, error) {
	config, cerr := NewConfig(configpath)
	if cerr != nil {
		logrus.Errorf("[ctx][NewCtx] %+v\n", cerr)
		return nil, cerr
	}

	logger, lerr := NewLogger(config)
	if lerr != nil {
		logrus.Errorf("[ctx][NewCtx] %+v\n", lerr)
		return nil, lerr
	}

	return &Ctx{Config: config, Logger: logger}, nil
}

// Close ...
func (c *Ctx) Close() error {
	c.Logger.logfile.Close()
	return nil
}
