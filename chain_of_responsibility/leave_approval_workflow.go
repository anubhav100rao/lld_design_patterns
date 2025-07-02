package chain_of_responsibility

import "fmt"

// Request
type LeaveRequest struct {
	Employee string
	Days     int
	Reason   string
}

// Handler interface
type Approver interface {
	SetNext(Approver)
	Approve(req *LeaveRequest)
}

// Base Handler
type BaseApprover struct {
	next Approver
}

func (b *BaseApprover) SetNext(next Approver) {
	b.next = next
}

func (b *BaseApprover) PassToNext(req *LeaveRequest) {
	if b.next != nil {
		b.next.Approve(req)
	} else {
		fmt.Println("Request reached end of chain; no one to handle it.")
	}
}

// Concrete Handlers

type Manager struct{ BaseApprover }

func (m *Manager) Approve(req *LeaveRequest) {
	if req.Days <= 2 {
		fmt.Printf("Manager approved %d days leave for %s\n", req.Days, req.Employee)
	} else {
		fmt.Println("Manager passes to Director")
		m.PassToNext(req)
	}
}

type Director struct{ BaseApprover }

func (d *Director) Approve(req *LeaveRequest) {
	if req.Days <= 5 {
		fmt.Printf("Director approved %d days leave for %s\n", req.Days, req.Employee)
	} else {
		fmt.Println("Director passes to CEO")
		d.PassToNext(req)
	}
}

type CEO struct{ BaseApprover }

func (c *CEO) Approve(req *LeaveRequest) {
	if req.Days <= 10 {
		fmt.Printf("CEO approved %d days leave for %s\n", req.Days, req.Employee)
	} else {
		fmt.Printf("Leave request for %d days exceeds limit; rejected\n", req.Days)
	}
}

// Client setup

func RunLeaveApprovalWorkflow() {
	// Instantiate handlers
	mgr := &Manager{}
	dir := &Director{}
	ceo := &CEO{}

	// Build the chain: Manager → Director → CEO
	mgr.SetNext(dir)
	dir.SetNext(ceo)

	// Test requests
	requests := []*LeaveRequest{
		{"Alice", 1, "Personal"},
		{"Bob", 4, "Vacation"},
		{"Carol", 7, "Medical"},
		{"Dave", 12, "Sabbatical"},
	}
	for _, req := range requests {
		fmt.Printf("\nRequesting %d days leave for %s:\n", req.Days, req.Employee)
		mgr.Approve(req)
	}
}
