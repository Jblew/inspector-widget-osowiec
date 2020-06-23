<template>
  <div class="osowiec-log-ui">
    <div>{{ lastLogTimeFormatted }}</div>
    <pre>
      {{ lastLogContents }}
    </pre>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch, Prop } from "vue-property-decorator";
import { StatefulResource, Resource } from "vue-stateful-resource";
import { listenForLogs } from "./listenForLogs";
import { LogEntry } from "../types";

@Component
export default class OsowiecLogUi extends Vue {
  @Prop({ required: true })
  logs!: LogEntry[];

  get lastLog(): LogEntry | undefined {
    return this.logs[0];
  }

  get lastLogContents(): string {
    return this.lastLog ? this.lastLog.contents : "";
  }

  get lastLogTimeFormatted(): string {
    if (this.lastLog) {
      return new Date(this.lastLog.timestamp).toISOString();
    } else {
      return "(no log)";
    }
  }
}
</script>
<style scoped>
.osowiec-log-ui {
  width: 100%;
  min-height: 40rem;
  overflow: scroll;
}
</style>
