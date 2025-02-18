import json
from kafka import KafkaProducer

KAFKA_BROKER = "kafka:9092"

def get_kafka_producer() -> KafkaProducer:
    producer = KafkaProducer(bootstrap_servers=[KAFKA_BROKER],
                             value_serializer=lambda v: json.dumps(v).encode('utf-8'))
    return producer

