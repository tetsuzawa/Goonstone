export const state = () => ({
  user: null
});

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
    const reseponse = await this.$axios.$post('/logout', data)
    context.commit('setUser', null)
  }
};

