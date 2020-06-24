<template>
  <div>
    <div class="osowiec-log-ui-status">
      Osowiec git backup log:
      {{ lastLogTimeFormatted }}
    </div>
    <div class="osowiec-log-ui" ref="logContainer">
      <pre>
      {{ lastLogContents }}
    </pre
      >
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Vue, Watch, Prop } from 'vue-property-decorator';
import { StatefulResource, Resource } from 'vue-stateful-resource';
import { listenForLogs } from './listenForLogs';
import { LogEntry } from '../types';

@Component
export default class OsowiecLogUI extends Vue {
  @Prop({ required: true })
  logs!: LogEntry[];

  get lastLog(): LogEntry | undefined {
    return this.logs[0];
  }

  get lastLogContents(): string {
    return this.lastLog ? this.lastLog.contents : '';
  }

  get lastLogTimeFormatted(): string {
    if (this.lastLog) {
      return new Date(this.lastLog.timestamp).toISOString();
    }
    return '(no log)';
  }

  mounted() {
    this.onLogsUpdated();
  }

  @Watch('logs')
  onLogsUpdated() {
    this.scrollBottom();
  }

  private scrollBottom() {
    const containerElem: any = this.$refs.logContainer;
    containerElem.scrollTop = containerElem.scrollHeight;
  }
}
</script>
<style scoped>
.osowiec-log-ui {
  width: 100%;
  height: 18rem;
  overflow: scroll;
  scroll-behavior: smooth;
}

.osowiec-log-ui-status {
  width: 100%;
  text-align: center;
  text-decoration: underline;
  color: #bbb;
}
</style>
