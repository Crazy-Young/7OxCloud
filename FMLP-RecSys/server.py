import io
import grpc
import transport_pb2 as pb2
import transport_pb2_grpc as pb2_grpc
from concurrent import futures
import time
import socket
from google.protobuf.empty_pb2 import Empty
from main import main2
from test import main


# pip install protobuf
# pip install types-protobuf

class TransportServiceServicer(pb2_grpc.TransportServiceServicer):
    def SendCSVFile(self, request, context):
        bytes_io = io.BytesIO(request.csv)
        with open(f"data/video-log.csv", "wb") as f:
            f.write(bytes_io.getvalue())
        main2()
        return Empty()

    def GetRecommendedFeed(self, request, context):
        print(request.userId)
        # if(request.userId==0):
        #     # data_file2 = args.data_dir + args.data_name + "-sorted-mapped" + '.json'
        #     # mydict = load(data_file2)
        rec_seq, flag = main(request.userId)

        return pb2.GetRecommendedFeedResponse(videoId=rec_seq[0])


def serve(port: int):
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb2_grpc.add_TransportServiceServicer_to_server(TransportServiceServicer(), server)
    server.add_insecure_port(f"[::]:{port}")
    server.start()
    print(f"Server started at {get_ipv4_address()}:{port}")
    try:
        while 1:
            time.sleep(600)
    except KeyboardInterrupt:
        server.stop(0)


def get_ipv4_address():
    s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    try:
        # doesn't even have to be reachable
        s.connect(('10.255.255.255', 1))
        IP = s.getsockname()[0]
    except Exception:
        IP = '127.0.0.1'
    finally:
        s.close()
    return IP


if __name__ == '__main__':
    serve(8976)
