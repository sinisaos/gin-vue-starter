import axios from 'axios'


const state = {
  user: undefined,
  tasks: undefined,
}

const getters = {
  isAuthenticated: state => !!state.user,
  stateUser: state => state.user,
}

const actions = {
  async loginUser(context, user) {
    const response = await axios.post('accounts/login', user)
    axios.defaults.headers.common['Authorization'] = `Bearer ${response.data.access_token}`
    context.dispatch('userProfile', response.data)
  },
  async userProfile(context) {
    const response = await axios.get('accounts/profile')
    context.commit('setUser', response.data.user)
  },
  async deleteUser() {
    await axios.delete(`accounts/delete`)
  },
  async logoutUser(context) {
    await axios.post('accounts/logout')
    context.commit('logout')
  },
  async userTasks(context) {
    const response = await axios.get('accounts/profile')
    context.commit('setTasks', response.data.user)
  },
}

const mutations = {
  setUser(state, username) {
    state.user = username
  },
  logout(state, user) {
    state.user = user
  },
}

export default {
  state,
  getters,
  actions,
  mutations
}
