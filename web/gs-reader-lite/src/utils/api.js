import { get, post } from './request'

export const LOGIN = (params) => post('/v1/api/user/login', params)
export const REGISTER = (params) => post('/v1/api/user/register', params)
export const GET_LATEST_FEED_ITEMS = (params) => get('/v1/api/feed/latest', params)
export const SEARCH_FEED_ITEMS = (params) => get('/v1/api/feed/search', params)
