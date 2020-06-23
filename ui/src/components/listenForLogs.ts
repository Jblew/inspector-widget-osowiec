import { Resource } from 'vue-stateful-resource';
import firebase from 'firebase/app';
import { listenForSnapshots } from './listenForSnapshots';
import { LogEntry } from '../types';
import { projectConfig } from '../project.config';

export const listenForLogs = (cb: (res: Resource<LogEntry[]>) => void) =>
  listenForSnapshots(getQuery(), processDocuments, cb);

function getQuery() {
  return getCollection()
    .limit(20)
    .orderBy('timestamp', 'desc');
}

function processDocuments(docs: firebase.firestore.QueryDocumentSnapshot[]) {
  return docs.map(doc => doc.data() as LogEntry);
}

function getCollection() {
  return firebase.firestore().collection(projectConfig.logCollection);
}
