// Copyright (C) 2019, Xiongfa Li.
// All right reserved.
// @author xiongfa.li
// @version V1.0
// Description:

package gobatis

import (
	"github.com/xfali/gobatis/datasource"
	"github.com/xfali/gobatis/factory"
	"github.com/xfali/gobatis/logging"
	"time"
)

type FacOpt func(f *factory.DefaultFactory)

func NewFactory(opts ...FacOpt) factory.Factory {
	f, _ := CreateFactory(opts...)
	return f
}

func CreateFactory(opts ...FacOpt) (factory.Factory, error) {
	f := &factory.DefaultFactory{
		Log: logging.DefaultLogf,
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			opt(f)
		}
	}

	err := f.Open(f.DataSource)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func SetMaxConn(v int) FacOpt {
	return func(f *factory.DefaultFactory) {
		f.MaxConn = v
	}
}

func SetMaxIdleConn(v int) FacOpt {
	return func(f *factory.DefaultFactory) {
		f.MaxIdleConn = v
	}
}

func SetConnMaxLifetime(v time.Duration) FacOpt {
	return func(f *factory.DefaultFactory) {
		f.ConnMaxLifetime = v
	}
}

func SetLog(v logging.LogFunc) FacOpt {
	return func(f *factory.DefaultFactory) {
		f.Log = v
	}
}

func SetDataSource(v datasource.DataSource) FacOpt {
	return func(f *factory.DefaultFactory) {
		f.WithLock(func(fac *factory.DefaultFactory) {
			fac.DataSource = v
		})
	}
}
