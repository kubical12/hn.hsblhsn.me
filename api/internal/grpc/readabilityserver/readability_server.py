import logging
import grpc
from os import environ
from concurrent.futures import ThreadPoolExecutor
from readability_pb2 import GetReadableDocumentRequest, GetReadableDocumentResponse, GetReadinessInfoResponse
from readability_pb2_grpc import ReadabilityServicer, add_ReadabilityServicer_to_server


from readability import Document


class ReadabilityServicerImpl(ReadabilityServicer):
    def GetReadableDocument(self, request: GetReadableDocumentRequest, context):
        logging.info(f'GetReadableDocument called for {request.identifier}')
        doc = Document(request.html)
        return GetReadableDocumentResponse(title=doc.short_title(), body=doc.summary())

    def GetReadinessInfo(self, request, context):
        logging.info(f'GetReadinessInfo called for {request.identifier}')
        return GetReadinessInfoResponse(ready=True)

if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    server = grpc.server(ThreadPoolExecutor(max_workers=10))
    add_ReadabilityServicer_to_server(ReadabilityServicerImpl(), server)
    listen_addr = environ.get('LISTEN_ADDRESS', "localhost:9595")
    server.add_insecure_port(listen_addr)
    server.start()
    logging.info('readability server ready on %r', listen_addr)
    try:
        server.wait_for_termination()
    except KeyboardInterrupt:
        logging.info('got keyboard interrupt')
    except Exception as e:
        logging.info('got unknown exception %r', e)
    finally:
        server.stop(True)
        logging.info('readability server stopped')
