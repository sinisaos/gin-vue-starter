import axios from 'axios'

const state = {
  tasks: undefined,
  task: undefined
}

const getters = {
  stateTasks: state => state.tasks,
  stateTask: state => state.task,
}

const actions = {
  async createTask(context, task) {
    await axios.post('tasks/', task)
    context.dispatch('userTasks')
  },
  async getTasks(context) {
    const response = await axios.get('accounts/profile')
    console.log(response.data.user)
    context.commit('setTasks', response.data.user)
  },
  async singleTask(context, id) {
    const response = await axios.get(`tasks/${id}`)
    context.commit('setTask', response.data)
  },
  async updateTask(context, task) {
    await axios.patch(`tasks/${task.id}`, task)
    context.dispatch('userTasks')
  },
  async deleteTask(context, id) {
    await axios.delete(`tasks/${id}`)
    context.dispatch('userTasks')
  }
}

const mutations = {
  setTasks(state, tasks) {
    state.tasks = tasks
  },
  setTask(state, task) {
    state.task = task
  },
};

export default {
  state,
  getters,
  actions,
  mutations
}
