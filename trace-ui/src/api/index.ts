import httpClient from '@/utils/http';
import base from '@/api/base';
import type {AppConnRes, AppRes, TraceListRes, TraceSpan} from '@/types';

export const listTraceApi = (data: any) => {
  return httpClient.post<TraceListRes>(base.LIST_TRACE_URL, data);
};

export const listAppApi = (data: any) => {
  return httpClient.post<AppRes[]>(base.GET_APP_URL, data);
};

export const treeTraceApi = (data: any) => {
  return httpClient.post<TraceSpan>(base.TREE_TRACE_URL, data);
};

export const getSpanInfo = (data: any) => {
  return httpClient.post<TraceSpan>(base.GET_SPAN_URL, data);
};

export const listAppConn = (data: any) => {
  return httpClient.post<{
    list: AppConnRes[]
  }>(base.GET_CONN_URL, data);
};



