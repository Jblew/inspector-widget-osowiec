<template>
  <stateful-resource :resource="logsResource">
    <osowiec-log-ui :logs="logs" />
  </stateful-resource>
</template>

<script lang="ts">
import { Component, Vue, Watch } from "vue-property-decorator";
import { StatefulResource, Resource } from "vue-stateful-resource";
import { listenForLogs } from "./listenForLogs";
import { LogEntry } from "../types";
import OsowiecLogUi from "./OsowiecLogUi.vue";

@Component({
  components: { StatefulResource }
})
export default class OsowiecLog extends Vue {
  logsResource: Resource<LogEntry[]> = Resource.empty();

  get logs(): LogEntry[] {
    const raw: LogEntry[] = [...(this.logsResource.result || [])];
    // tslint:disable no-console
    /* eslint-disable */
    console.log(raw);
    return raw.sort((a, b) => a.timestamp - b.timestamp);
  }

  beforeMount() {
    listenForLogs(res => {
      this.logsResource = res;
    });
  }
}
</script>
