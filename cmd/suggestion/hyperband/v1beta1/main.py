import grpc
import time
import logging
from concurrent import futures

from pkg.apis.manager.v1beta1.python import api_pb2_grpc
from pkg.apis.manager.health.python import health_pb2_grpc
from pkg.suggestion.v1beta1.hyperband.service import HyperbandService


_ONE_DAY_IN_SECONDS = 60 * 60 * 24
DEFAULT_PORT = "0.0.0.0:6789"

logger = logging.getLogger()
logging.basicConfig(level=logging.INFO)


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    service = HyperbandService()
    api_pb2_grpc.add_SuggestionServicer_to_server(service, server)
    health_pb2_grpc.add_HealthServicer_to_server(service, server)

    server.add_insecure_port(DEFAULT_PORT)
    logger.info("Start Hyperband service at address {}".format(DEFAULT_PORT))
    server.start()
    try:
        while True:
            time.sleep(_ONE_DAY_IN_SECONDS)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == "__main__":
    serve()
