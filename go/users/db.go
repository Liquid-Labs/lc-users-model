package users

import (
  "github.com/go-pg/pg/orm"

  . "github.com/Liquid-Labs/terror/go/terror"
  . "github.com/Liquid-Labs/lc-entities-model/go/entities"
)
/*
func (u *User) hideDeletedAt() {
  u.deletedAt = u.DeletedAt
}

func (u *User) restoreDeletedAt() {
  u.DeletedAt = u.deletedAt
}*/

// ModelEntity provides a(n initially empty) Entity receiver and base query.
func ModelUser(db orm.DB) (*User, *orm.Query) {
  u := &User{}
  q := db.Model(u)

  return u, q
}

// Create creates (or inserts) a new User record into the DB. As Users are logically abstract, one would typically only call this as part of another items create sequence.
func (u *User) Create(db orm.DB) Terror {
  if err := (&u.Subject.Entity).Create(db); err != nil {
    return err
  } else {
    qs := db.Model((&u.Subject)).ExcludeColumn(EntityFields...)
    if _, err := qs.Insert(); err != nil {
      return ServerError(`There was a problem creating the subject record.`, err)
    } else {
      qu := db.Model(u).ExcludeColumn(EntityFields...)
      if _, err := qu.Insert(); err != nil {
        return ServerError(`There was a problem creating the user record.`, err)
      } else {
        return nil
      }
    }
  }
}

var updateExcludes = make([]string, len(EntityFields))
func init() {
  copy(updateExcludes, EntityFields)
  updateExcludes = append(updateExcludes, "id")
}
// Updates a User record in the DB. As Users are logically abstract, one would typically only call this as part of another items update sequence.
func (u *User) Update(db orm.DB) Terror {
  // u.hideDeletedAt()
  // defer u.restoreDeletedAt()
  if err := (&u.Subject.Entity).Update(db); err != nil {
    return err
  } else {
    /* So, there's really nothing on 'subjects' to update and when all columns
       are excluded, go-pg treats it the same as if nothing was excluded, which
       leads to exceptions when it tries to update the entity fields.
    qs := db.Model((&u.Subject)).
      ExcludeColumn(updateExcludes...).
      Where(`subject.id=?id`)
    qs.GetModel().Table().SoftDeleteField = nil
    if _, err := qs.Update(); err != nil {
      return ServerError(`There was a problem updating the subject record.`, err)
    } else { */
      qu := db.Model(u).
        ExcludeColumn(updateExcludes...).
        Where(`"user".id=?id`)
      qu.GetModel().Table().SoftDeleteField = nil
      if _, err := qu.Update(); err != nil {
        return ServerError(`There was a problem updating the user record.`, err)
      } else {
        return nil
      }
    // }
  }
}

// Archive updates a User record in the DB. As Users are logically abstract, one would typically only call this as part of another items archive sequence.
func (u *User) Archive(db orm.DB) Terror {
  return (&u.Entity).Archive(db)
}