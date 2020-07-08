package main
import (
  "log"
  "io"
  pb "github.com/grpc-streaming-demo/protobuf_kedro"
  "fmt"

  //   "gitlab.com/pantomath-io/demo-grpc/api"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
)
func main() {
  var conn *grpc.ClientConn
  conn, err := grpc.Dial(":50051", grpc.WithInsecure())
  if err != nil {
    log.Fatalf("did not connect: %s", err)
  }
  defer conn.Close()
  c := pb.NewKedroClient(conn)
  fmt.Println(c)
  response, err := c.ListPipelines(context.Background(), &pb.PipelineParams{})

  if err != nil {
    log.Fatalf("Error when calling SayHello: %s", err)
  }
  log.Printf("Response from server: %s", response)

  run, err := c.Run(context.Background(), &pb.RunParams{})
  if err != nil {
    log.Fatalf("%v.Run(_) = _, %v", c, err)
  }

  log.Println(run)
  statusStream, err := c.Status(context.Background(), &pb.RunId{RunId: run.RunId})

  for {
    status, err := statusStream.Recv()
    if err == io.EOF {
      break
    }
    if err != nil {
      log.Fatalf("%v.Status(_) = _, %v", c, err)
    }
    log.Println(status.GetEvents())
    if status.GetRunStatus() == "Completed" {
      log.Println(status.GetRunStatus())
      log.Println(status.GetSuccess())
    }
  }
}

// func (s *routeGuideServer) ListFeatures(rect *pb.Rectangle, stream pb.RouteGuide_ListFeaturesServer) error {
// 	for _, feature := range s.savedFeatures {
// 		if inRange(feature.Location, rect) {
// 			if err := stream.Send(feature); err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return nil
// }

// // rect := &pb.Rectangle{ ... }  // initialize a pb.Rectangle
// stream, err := client.ListFeatures(context.Background(), rect)
// if err != nil {
//     ...
// }
// for {
//     feature, err := stream.Recv()
//     if err == io.EOF {
//         break
//     }
//     if err != nil {
//         log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
//     }
//     log.Println(feature)
// }


