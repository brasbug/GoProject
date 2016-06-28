package main

import (
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

type Person struct {
	Name  string
	Phone string
	id string
}

type User struct {
	Id_       bson.ObjectId `bson:"_id"`
	Name      string        `bson:"name"`
	Age       int           `bson:"age"`
	JoinedAt   time.Time     `bson:"joined_at"`
	Interests []string      `bson:"interests"`
}



func main() {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//c := session.DB("test").C("people")
	//err = c.Insert(&Person{"Ale", "+55 53 8116 9639","1001"},
	//    &Person{"Cla", "+55 53 8402 8510","1002"})
	//if err != nil {
	//    panic(err)
	//}
	//
	////err = c.up
	//result := Person{}
	//err = c.Find(bson.M{"id": "1002"}).One(&result)
	//if err != nil {
	//    panic(err)
	//}
	//
	//fmt.Println("Phone:", result.Phone)


	c := session.DB("test").C("user")

	//err = c.Insert(&User{
	//    Id_:       bson.NewObjectId(),
	//    Name:      "Jimmy Kuu",
	//    Age:       33,
	//    JoinedAt:  time.Now(),
	//    Interests: []string{"Develop", "Movie"},
	//})
	//
	//if err != nil {
	//    panic(err)
	//}
	//
	//err = c.Insert(&User{
	//    Id_:       bson.NewObjectId(),
	//    Name:      "Tracy Yu",
	//    Age:       31,
	//    JoinedAt:  time.Now(),
	//    Interests: []string{"Shoping", "TV"},
	//})

	//if err != nil {
	//    panic(err)
	//}

	//var users []User
	//c.Find(nil).All(&users)
	//fmt.Println(users)

	//objectId := bson.ObjectIdHex("56fc874ae92e6b0419000001")
	//
	//
	//user := new(User)
	////c.Find(bson.M{"name": "Tracy Yu"}).One(&user)
	//c.FindId(objectId).One(&user)
	//
	//fmt.Println(user)

	users  := []User{}
	//c.Find(bson.M{"name": "Jimmy Kuu"}).All(&users)
	//c.Find(bson.M{"name": bson.M{"$ne": "Jimmy Kuu"}}).All(&users)
	//c.Find(bson.M{"name": bson.M{"$in": []string{"Jimmy Kuu", "Tracy Yu"}}}).All(&users)


	//err = c.Find(bson.M{"$or": []bson.M{bson.M{"name": "Jimmy Kuu"}, bson.M{"age": 31}}}).All(&users)

	//c.Update(bson.M{"_id": objectId},
	//    bson.M{"$set": bson.M{
	//        "name": "Jimmy Gu",
	//        "age":  34,
	//    }})

	//c.Update(bson.M{"_id": objectId},
	//    bson.M{"$inc": bson.M{
	//        "age": -1,
	//    }})

	//c.Update(bson.M{"_id": objectId},
	//    bson.M{"$push": bson.M{
	//        "interests": "Golang",
	//    }})
	//
	//c.Update(bson.M{"_id": objectId},
	//    bson.M{"$pull": bson.M{
	//        "interests": "Golang",
	//    }})

	//c.Remove(bson.M{"name": "Tracy Yu"})
	err = c.Find(bson.M{"name": "Tracy Yu"}).All(&users)

	if err != nil{
		panic(err)
	}
	for i := 0; i < len(users);i ++  {
		user1 := users[i];
		fmt.Println(user1)
	}

}