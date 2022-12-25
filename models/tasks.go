package models

import (
	"encoding/binary"
	"fmt"

	bolt "go.etcd.io/bbolt"
)

type TaskService struct {
	DB *bolt.DB
}

func (ts TaskService) ListTasks() error {
	err := ts.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("tasks"))
		if b == nil {
			fmt.Println("Tasks have never been added!")
			return nil
		}
		c := b.Cursor()
		id, task := c.First()
		if id == nil {
			fmt.Println("List is empty!")
			return nil
		}
		for ; id != nil; id, task = c.Next() {
			intID := btoi(id)
			fmt.Printf("%d. %s\n", intID, task)
		}
		return nil
	})
	return err
}

func (ts TaskService) AddTask(task string) error {
	err := ts.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return fmt.Errorf("add tasks: %w", err)
		}
		lastElemKey, _ := b.Cursor().Last()
		if lastElemKey == nil {
			lastElemKey = make([]byte, 8)
		}
		id := btoi(lastElemKey) + 1
		return b.Put(itob(id), []byte(task))
	})
	if err == nil {
		fmt.Println("Added: ", task)
	}
	return err
}

func (ts TaskService) DoTask(taskID int) error {
	err := ts.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("tasks"))
		if err != nil {
			return fmt.Errorf("list tasks: %w", err)
		}
		id := itob(uint64(taskID))
		task := b.Get(id)
		if task == nil {
			fmt.Println("There is now such task!")
			return nil
		}
		err = b.Delete(id)
		if err != nil {
			return fmt.Errorf("do task: %w", err)
		}
		fmt.Println("Done: ", string(task))
		return nil
	})
	return err
}

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

// boti makes reverse conversion to itob
func btoi(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}
