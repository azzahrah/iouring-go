// +build linux

package iouring

// See uapi/linux/io_uring.h

// Opcode is an opcode for the ring.
type Opcode uint8

const (
	// SetupSyscall defines the syscall number for io_uring_setup.
	SetupSyscall = 425
	// EnterSyscall defines the syscall number for io_uring_enter.
	EnterSyscall = 426
	// RegisterSyscall defines the syscall number for io_uring_register.
	RegisterSyscall = 427
)

const (

	/*
	 * io_uring_params->features flags
	 */
	FeatSingleMmap     = (1 << 0)
	FeatNoDrop         = (1 << 1)
	FeatSubMitStable   = (1 << 2)
	FeatRwCurPos       = (1 << 3)
	FeatCurPersonality = (1 << 4)

	/*
	 * sqe->flags
	 */

	// SqeFixedFile use fixed fileset
	SqeFixedFile uint = (1 << 0)
	// SqeIODrain issue after inflight IO
	SqeIODrain uint = (1 << 1)

	/*
	 * io_uring_setup() flags
	 */

	// SetupIOPoll io_context is polled
	SetupIOPoll uint = (1 << 0)
	// SetupSQPoll SQ poll thread
	SetupSQPoll uint = (1 << 1)
	// SetupSQAFF sq_thread_cpu is valid
	SetupSQAFF uint = (1 << 2)

	Nop Opcode = iota
	Readv
	Writev
	Fsync
	ReadFixed
	WriteFixed
	PollAdd
	PollRemove
	SyncFileRange
	SendMsg
	RecvMsg
	Timeout
	TimeoutRemove
	Accept
	AsyncCancel
	LinkTimeout
	Connect
	Fallocate
	OpenAt
	Close
	FilesUpdate
	Statx
	Read
	Write
	Fadvise
	Madvise
	Send
	Recv
	Openat2
	EpollCtl
	Splice
	ProvideBuffers
	RemoveBuffers

	OpSupported = (1 << 0)

	/*
	 * sqe->fsync_flags
	 */

	// FsyncDatasync ...
	FsyncDatasync uint = (1 << 0)

	/*
	 * Magic offsets for the application to mmap the data it needs
	 */

	// SqRingOffset is the offset of the submission queue.
	SqRingOffset uint64 = 0
	// CqRingOffset is the offset of the completion queue.
	CqRingOffset uint64 = 0x8000000
	// SqeSOffset is the offset of the submission queue entries.
	SqeSOffset uint64 = 0x10000000

	/*
	 * sq_ring->flags
	 */

	// SqNeedWakeup needs io_uring_enter wakeup
	SqNeedWakeup uint = (1 << 0)

	/*
	 * io_uring_enter(2) flags
	 */

	// EnterGetEvents ...
	EnterGetEvents uint = (1 << 0)
	// EnterSQWakeup ...
	EnterSQWakeup uint = (1 << 1)

	/*
	 * io_uring_register(2) opcodes and arguments
	 */

	RegisterBuffers   = 0
	UnregisterBuffers = 1
	RegisterFiles     = 2
	UnregisterFiles   = 3
	RegisterEventfd   = 4
	UnregisteREventfd = 5
)
