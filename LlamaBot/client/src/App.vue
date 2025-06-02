<template>
  <div class="mx-auto max-w-100 py-3">
    <ul v-if="messages.length">
      <li
        v-for="message in messages"
        class="border border-green-300 p-3 mb-5 w-70 rounded"
      >
        <strong>{{ message.isAgent ? 'Agent' : 'You' }}:</strong>
        {{ message.content }}
      </li>
    </ul>

    <p v-if="submitting" class="border border-green-300 p-3 mb-5 w-70 rounded">
      <strong>Agent:</strong> writing...
    </p>

    <form @submit.prevent="submit">
      <textarea
        v-model="form.content"
        class="p-3 border border-blue-500 rounded"
      />
      <button
        type="submit"
        class="block p-3 rounded border border-blue-400 w-30"
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
      content: null
    },
    messages: [],
    submitting: false
  }),

  methods: {
    async submit() {
      this.submitting = true

      this.addMessage(this.form.content)

      const { data } = await axios.post(`${BASE_URL}/send_question`, {
        content: this.form.content
      })

      this.addMessage(data.content, true)

      this.submitting = false
      this.clearForm()
    },

    addMessage(content, isAgent = false) {
      this.messages.push({
        content,
        isAgent
      })
    },

    clearForm() {
      this.form.content = null
    }
  }
}
</script>
