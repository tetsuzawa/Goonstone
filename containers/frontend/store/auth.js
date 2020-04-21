import {STATUS_CREATED, STATUS_OK} from "../static/js/errorCodes";

export const state = () => ({
  user: null,
  apiStatusHasError: null
});

export const getters = {
  hasLoggedIn: state => !!state.user,
  username: state => state.user ? state.user.name : ''
};

export const mutations = {
  setUser(state, user) {
    state.user = user
  },
  setAPIStatus(state, status) {
    state.apiStatusHasError = status
  }
};

export const actions = {
  async register(context, data) {
    const response = await this.$axios.post('/register', data)
      .catch(error => error.response || error);
    context.commit('setUser', response.data.user)
    if (response.status === STATUS_CREATED) {
      context.commit('setAPIStatus', true);
      context.commit('setUser', response.data.user);
      return false
    }
    context.commit('setAPIStatus', false);
    context.commit('error/setCode', response.status, {root: true})
  },
  async login(context, data) {
    context.commit('setAPIStatus', null);
    const response = await this.$axios.post('/login', data)
      .catch(error => error.response || error);
    if (response.status === STATUS_OK) {
      context.commit('setAPIStatus', true);
      context.commit('setUser', response.data.user);
      return false
    }
    context.commit('setAPIStatus', false);
    context.commit('error/setCode', response.status, {root: true})
  },
  async logout(context) {
    const response = await this.$axios.post('/logout', null)
      .catch(error => error.response || error);
    if (response.status !== STATUS_OK) {
      context.commit('error/setCode', response.status, {root: true})
    }
    context.commit('setUser', null)
  },
  async currentUser(context) {
    const response = await this.$axios.get('/user')
      .catch(error => error.response || error);
    if (response.status === STATUS_OK) {
      const user = response.data.user || null;
      context.commit('setAPIStatus', true);
      context.commit('setUser', user);
      return false
    }
    context.commit('setAPIStatus', false);
    context.commit('error/setCode', response.status, {root: true})
  }
};

