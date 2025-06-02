<template>
  <div class="mx-auto max-w-[600px] py-3 border border-gray-200 p-6">
    <ul
      class="w-full max-h-[900px] h-[900px] overflow-y-auto flex flex-col"
      ref="messages"
    >
      <li
        v-for="message in messages"
        class="border p-3 mb-5 w-70 rounded w-full max-w-[360px]"
        :class="{
          'ml-auto border-blue-300': message.isAgent,
          'border-green-300': !message.isAgent
        }"
      >
        <strong>{{ message.isAgent ? 'Agent' : 'You' }}:</strong>
        {{ message.content }}
      </li>

      <p
        v-if="submitting"
        class="border border-green-300 p-3 mb-5 w-70 rounded w-full max-w-[360px] ml-auto"
      >
        <strong>Agent:</strong> Typing...
      </p>
    </ul>

    <form @submit.prevent="submit">
      <textarea
        v-model.trim="form.content"
        class="p-3 border border-blue-500 rounded w-full"
      />
      <button
        type="submit"
        class="block p-3 rounded border border-blue-400 w-30 w-full"
        :disabled="submitting"
      >
        Send
      </button>
    </form>
  </div>
</template>

<script>
import axios from 'axios'

const BASE_URL = 'http://localhost:8000'

export default {
  data: () => ({
    form: {
      content: ''
    },
    messages: [],
    submitting: false
  }),

  methods: {
    async submit() {
      const content = this.form.content

      if (content == '') {
        return false
      }
      this.clearForm()

      this.addMessage(content)
      await this.sendMessage(content)
    },

    addMessage(content, isAgent = false) {
      this.messages.push({
        content,
        isAgent
      })

      this.$nextTick(() => {
        const el = this.$refs.messages
        el.scrollTop = el.scrollHeight
      })
    },

    async sendMessage(content) {
      this.submitting = true

      const { data } = await axios.post(`${BASE_URL}/chat`, {
        content
      })

      this.addMessage(data.content, true)

      this.submitting = false
    },

    clearForm() {
      this.form.content = null
    }
  },

  created() {
    this.sendMessage('START_CHAT')
  }
}
</script>
