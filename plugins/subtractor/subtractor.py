from concurrent import futures
import grpc
import sys
import time

from grpc_health.v1.health import HealthServicer
from grpc_health.v1 import health_pb2, health_pb2_grpc
import grpc_stdio_pb2_grpc
from grpc_stdio import Logger, StdioService

from mather_pb2 import MathResponse
import mather_pb2_grpc

logger = Logger()
log = logger.log


class SubtractorServicer(mather_pb2_grpc.MatherServicer):
    def DoMath(self, request, context):
        return MathResponse(result=request.x-request.y)


def serve():
    # We need to build a health service to work with go-plugin
    health = HealthServicer()
    health.set(
        "subtractor", health_pb2.HealthCheckResponse.ServingStatus.Value('SERVING'))

    # Make the server
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    subtractor = SubtractorServicer()

    # Add our service
    mather_pb2_grpc.add_MatherServicer_to_server(subtractor, server)
    health_pb2_grpc.add_HealthServicer_to_server(health, server)
    grpc_stdio_pb2_grpc.add_GRPCStdioServicer_to_server(
        StdioService(logger), server)

    # Listen on a port
    server.add_insecure_port('127.0.0.1:51010')

    # Start
    server.start()
    # Output information
    print("1|1|tcp|127.0.0.1:51010|grpc")
    sys.stdout.flush()

    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
