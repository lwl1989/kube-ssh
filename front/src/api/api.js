import request from '@/utils/request'

export function clusterList() {
  return request({
    url: '/k8s/clusters',
    method: 'get'
  })
}

export function podList(query) {
  return request({
    url: '/k8s/workload/pods?id=' + query.id,
    method: 'get'
  })
}

export function signCluster(data) {
  return request({
    url: '/k8s/sign',
    method: 'post',
    data
  })
}

export function userList() {
  return request({
    url: '/users',
    method: 'get'
  })
}

export function managerList(query) {
  return request({
    url: '/managers?page=' + query.page + '&size=' + query.size,
    method: 'get'
  })
}

export function managerUpsert(data) {
  return request({
    url: '/manager/upsert',
    method: 'post',
    data
  })
}

export function managerStatus(data) {
  return request({
    url: '/manager/status',
    method: 'post',
    data
  })
}

export function whitesList(query) {
  return request({
    url: '/whites?page=' + query.page + '&size=' + query.size,
    method: 'get'
  })
}

export function whiteUpsert(data) {
  return request({
    url: '/white',
    method: 'post',
    data
  })
}

export function whiteStatus(data) {
  return request({
    url: '/white/status',
    method: 'post',
    data
  })
}

export function whiteDelete(data) {
  return request({
    url: '/white',
    method: 'delete',
    data
  })
}
