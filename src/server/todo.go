package server

import (
	"fmt"
	"log"
	"runtime"
	"runtime/debug"

	"golang.org/x/net/context"

	pb "datamodel/protobuf"
)

func (s *serverStruct) CreateDomain(ctx context.Context, in *pb.Domain) (*pb.Domain, error) {
	return in, nil
}

func (s *serverStruct) ListDomains(ctx context.Context, in *pb.Empty) (*pb.ListDomainsReply, error) {
	doms := &pb.ListDomainsReply{}
	return doms, nil
}

func (s *serverStruct) DeleteDomain(ctx context.Context, in *pb.Domain) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *serverStruct) GetDomain(ctx context.Context, in *pb.Domain) (*pb.Domain, error) {
	return in, nil
}

func (s *serverStruct) GetSnapshot(ctx context.Context, in *pb.GetSnapshotRequest) (*pb.GetSnapshotReply, error) {
	return nil, nil
}

func (s *serverStruct) CreateSnapshot(ctx context.Context, in *pb.CreateSnapshotRequest) (*pb.CreateSnapshotReply, error) {
	return nil, nil
}

func (s *serverStruct) CreateSketch(ctx context.Context, in *pb.Sketch) (*pb.Sketch, error) {
	return in, nil
}

func (s *serverStruct) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddReply, error) {
	m := &runtime.MemStats{}
	fmt.Println("========== Start")
	log.Println("# goroutines: ", runtime.NumGoroutine())
	runtime.ReadMemStats(m)
	log.Println("Memory Acquired: ", m.Sys)
	log.Println("Memory Used    : ", m.Alloc)
	fmt.Println("--- run GC")
	debug.FreeOSMemory()
	log.Println("# goroutines: ", runtime.NumGoroutine())
	runtime.ReadMemStats(m)
	log.Println("Memory Acquired: ", m.Sys)
	log.Println("Memory Used    : ", m.Alloc)
	fmt.Println("========== End")
	return &pb.AddReply{}, nil
}

func (s *serverStruct) GetMembership(ctx context.Context, in *pb.GetRequest) (*pb.GetMembershipReply, error) {
	reply := &pb.GetMembershipReply{}
	return reply, nil
}

func (s *serverStruct) GetFrequency(ctx context.Context, in *pb.GetRequest) (*pb.GetFrequencyReply, error) {
	reply := &pb.GetFrequencyReply{}
	return reply, nil
}

func (s *serverStruct) GetCardinality(ctx context.Context, in *pb.GetRequest) (*pb.GetCardinalityReply, error) {
	reply := &pb.GetCardinalityReply{}
	return reply, nil
}

func (s *serverStruct) GetRankings(ctx context.Context, in *pb.GetRequest) (*pb.GetRankingsReply, error) {
	reply := &pb.GetRankingsReply{}
	return reply, nil
}

func (s *serverStruct) DeleteSketch(ctx context.Context, in *pb.Sketch) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s *serverStruct) ListAll(ctx context.Context, in *pb.Empty) (*pb.ListReply, error) {
	filtered := &pb.ListReply{}
	return filtered, nil
}

func (s *serverStruct) GetSketch(ctx context.Context, in *pb.Sketch) (*pb.Sketch, error) {
	return in, nil
}

func (s *serverStruct) List(ctx context.Context, in *pb.ListRequest) (*pb.ListReply, error) {
	filtered := &pb.ListReply{}
	return filtered, nil
}
