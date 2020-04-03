export const state = () => ({
  user: null
});

export const getters = {
    hasLoggedIn: state => !!state.user,
    username: state => state.user ? state.user.name : ''
  }
;

export const mutations = {
  setUser(state, user) {
    state.user = user
  },
};

export const actions = {
  async register(context, data) {
    try {
      const response = await this.$axios.$post('/register', data);
      context.commit('setUser', response.user)
    } catch (error) {
      console.log("error: login: %O", error);
    }
  },
  async login(context, data) {
    try {
      const response = await this.$axios.$post('/login', data);
      context.commit('setUser', response.user)
    } catch (error) {
      console.log("error: login: %O", error);
    }
  },
  async logout(context) {
    try {
      const response = await this.$axios.$post('/logout', null)
      context.commit('setUser', null)
    } catch (error) {
      console.log("error: logout: %O", error);
    }
  },
  async currentUser(context) {
    try {
      const response = await this.$axios.$get('/user')
      const user = response.user || null
      context.commit('setUser', user)
    } catch (error) {
      console.log("error: logout: %O", error);
    }
  }
};

