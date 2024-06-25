import { QdrantClient } from '@qdrant/qdrant-js';

const client = new QdrantClient({url: 'http://127.0.0.1:6333'});

client.deleteCollection("global_patient_data");
