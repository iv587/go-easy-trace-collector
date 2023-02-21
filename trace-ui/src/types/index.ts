export interface TraceSearchForm {
  error: number,
  day?: string,
  operationKey?: string,
  applicationGroup?: string,
  application?: string
}

export interface TraceListRes {
  total: number,
  list: TraceSpan[]

}

export interface TraceSpan {
  operationName: string,
  startTime: number,
  finishTime: number,
  applicationGroup: string,
  appInstance: string,
  application: string,
  component: string
  elapsedTime: number,
  startTimeText: string
  tagType: string,
  children: TraceSpan[],
  deepth: number,
  traceId: string,
  tags: string,
  logDatas: string,
}

export interface AppRes {
  label: string,
  value: string,
  children: AppRes[]
}

export interface AppConnRes {
  addr: string;
  aliveTime: string;
  appGroup: string;
  appName: string;
  appStartTime: string;
  createTime: string;
  netInput: string;
}