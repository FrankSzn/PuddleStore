package puddlestore

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var Debug *log.Logger
var Out *log.Logger
var Error *log.Logger

// Initialize the loggers
func init() {
	Debug = log.New(ioutil.Discard, "", log.Ltime|log.Lshortfile)
	Out = log.New(os.Stdout, "", log.Ltime|log.Lshortfile)
	Error = log.New(os.Stdout, "ERROR: ", log.Ltime|log.Lshortfile)
}

// Turn debug on or off
func SetDebug(enabled bool) {
	if enabled {
		Debug = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		Debug = log.New(ioutil.Discard, "", log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func (r *PuddleNode) Out(formatString string, args ...interface{}) {
	Out.Output(2, fmt.Sprintf("%v", fmt.Sprintf(formatString, args...)))
}

func (r *PuddleNode) Debug(formatString string, args ...interface{}) {
	Debug.Output(2, fmt.Sprintf("%v", fmt.Sprintf(formatString, args...)))
}

func (r *PuddleNode) Error(formatString string, args ...interface{}) {
	Error.Output(2, fmt.Sprintf("%v", fmt.Sprintf(formatString, args...)))
}

/*
func (s NodeState) String() string {
	switch s {
	case FOLLOWER_STATE:
		return "follower"
	case CANDIDATE_STATE:
		return "candidate"
	case LEADER_STATE:
		return "leader"
	case JOIN_STATE:
		return "joining"
	default:
		return "unknown"
	}
}

func FsmCommandString(cmd FsmCommand) string {
	switch cmd {
	case HASH_CHAIN_ADD:
		return "hash-chain-add"
	case HASH_CHAIN_INIT:
		return "hash-chain-init"
	case CLIENT_REGISTRATION:
		return "client-registration"
	case INIT:
		return "init"
	case NOOP:
		return "noop"
	default:
		return "unknown"
	}
}

func (r *RaftNode) ShowState() {
	fmt.Printf("Current node state:\n")
	for i, otherNode := range r.GetOtherNodes() {
		fmt.Printf("%v - %v", i, otherNode)
		local := *r.GetLocalAddr()

		if local == otherNode {
			fmt.Printf(" (local node)")
		}
		if r.LeaderAddr != nil &&
			otherNode == *r.LeaderAddr {
			fmt.Printf(" (leader node)")
		}
		fmt.Printf("\n")

	}
	fmt.Printf("Current term: %v\n", r.GetCurrentTerm())
	fmt.Printf("Current state: %v\n", r.State)
	fmt.Printf("Current commit index: %v\n", r.commitIndex)
	fmt.Printf("Current next index: %v\n", r.nextIndex)
	fmt.Printf("Current match index: %v\n", r.matchIndex)

}

func (r *RaftNode) PrintLogCache() {
	fmt.Printf("Node %v LogCache:\n", r.Id)
	for _, entry := range r.logCache {
		fmt.Printf(" idx:%v, term:%v\n", entry.Index, entry.TermId)
	}
}*/
