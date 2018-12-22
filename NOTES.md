### API for pkg

us = UserPkg{
db = mysqldb,
inmemdb = inmem,
}
err = us.usecase.Store(entity.User{})
