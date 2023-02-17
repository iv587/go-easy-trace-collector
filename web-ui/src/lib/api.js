import { post } from './http'

export const getTraceListApi = data => post("/api/trace/list",data)

export const getTraceTreeApi = data => post("/api/trace/tree",data)

export const getSpanByIdApi = data => post("/api/trace/getSpanById",data)

export const getAppApi = data => post("/api/trace/getApp",data)


export const getConnectList = data => post("/api/connect/list",data)
