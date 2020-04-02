export const state = () => ({
  user: null
});

export const getters = {
  check: state => !!state.user,
  username: state => state.user ? state.user.name : ''
};

export const mutations = {
  setUser(state, user) {
    state.user = user
  },
};

export const actions = {
  async register(context, data) {
    try {
      const response = await this.$axios.$post('/register', data);
      context.commit('setUser', response.data)
    } catch (error) {
      console.log(error);
    }
  },
  async login(context, data) {
    try {
      const response = await this.$axios.$post('/login', data);
      context.commit('setUser', response.data)
    } catch (error) {
      console.log(error);
    }
  },
  async logout(context) {
    try {
      const response = await this.$axios.$post('/logout', null)
      context.commit('setUser', null)
    } catch (error) {
      console.log(error);
    }
  }
};

