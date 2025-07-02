package command

import "fmt"

// Command interface
type DBCommand interface {
	Execute() error
	Rollback() error
}

// Receiver (mock)
type Database struct{}

func (db *Database) Insert(table string, id int, data string) error {
	fmt.Printf("INSERT INTO %s VALUES (%d, %q)\n", table, id, data)
	return nil
}

func (db *Database) Delete(table string, id int) error {
	fmt.Printf("DELETE FROM %s WHERE id=%d\n", table, id)
	return nil
}

// Concrete Commands
type InsertCommand struct {
	db    *Database
	table string
	id    int
	data  string
}

func (c *InsertCommand) Execute() error {
	return c.db.Insert(c.table, c.id, c.data)
}
func (c *InsertCommand) Rollback() error {
	return c.db.Delete(c.table, c.id)
}

// Invoker
type TransactionManager struct {
	commands []DBCommand
}

func (tm *TransactionManager) Add(c DBCommand) {
	tm.commands = append(tm.commands, c)
}

func (tm *TransactionManager) Commit() {
	for _, c := range tm.commands {
		if err := c.Execute(); err != nil {
			fmt.Println("Error executing, rolling back")
			tm.Rollback()
			return
		}
	}
	fmt.Println("Transaction committed")
}

func (tm *TransactionManager) Rollback() {
	// undo in reverse order
	for i := len(tm.commands) - 1; i >= 0; i-- {
		tm.commands[i].Rollback()
	}
	fmt.Println("Transaction rolled back")
}

// Client
func RunDBTransactionCommandQueue() {
	db := &Database{}
	tm := &TransactionManager{}

	tm.Add(&InsertCommand{db, "users", 1, "Alice"})
	tm.Add(&InsertCommand{db, "users", 2, "Bob"})

	tm.Commit()
}
