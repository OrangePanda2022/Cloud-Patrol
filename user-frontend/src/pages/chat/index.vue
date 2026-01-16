<template>
  <div class="min-h-screen bg-base-200 flex flex-col items-center p-4 sm:p-8">
    <div
      class="card w-full max-w-4xl bg-base-100 shadow-md mb-4 overflow-hidden border border-base-300"
    >
      <div class="flex flex-col h-[70vh]">
        <div class="flex-1 overflow-y-auto p-4 space-y-4 scroll-smooth bg-dot-pattern">
          <div v-if="answer" class="chat chat-start">
            <div class="chat-image avatar">
              <div class="w-10 rounded-full bg-primary/10 p-1">
                <span class="text-xs flex justify-center mt-2">🤖</span>
              </div>
            </div>
            <div class="chat-bubble chat-bubble-primary leading-relaxed shadow-sm">
              <pre class="whitespace-pre-wrap font-sans">{{ answer }}</pre>
            </div>
          </div>
        </div>

        <div class="p-4 bg-base-100 border-base-200">
          <div class="p-4 bg-base-100 border-t border-base-200">
            <div class="join flex w-full">
              <textarea
                v-model="question"
                @keyup.enter.exact.prevent="startChat"
                class="textarea textarea-bordered join-item flex-1 min-h-[3rem] max-h-48 resize-none focus:outline-none"
                placeholder="在此输入您的问题..."
                rows="1"
              ></textarea>

              <button
                class="btn btn-primary join-item h-auto min-h-[3rem]"
                @click="startChat"
                :disabled="loading || !question.trim()"
              >
                <div v-if="loading" class="loading loading-spinner"></div>
                <div v-else class="flex items-center gap-2">
                  <span class="hidden sm:inline">发送</span>
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    class="w-5 h-5"
                  >
                    <path
                      stroke-linecap="round"
                      stroke-linejoin="round"
                      d="M6 12L3.269 3.126A59.768 59.768 0 0121.485 12 59.77 59.77 0 013.27 20.876L5.999 12zm0 0h7.5"
                    />
                  </svg>
                </div>
              </button>
            </div>
          </div>
          <p class="text-xs text-center opacity-40">AI 生成内容仅供参考</p>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* 让滚动条更精致 */
::-webkit-scrollbar {
  width: 5px;
}
::-webkit-scrollbar-thumb {
  background: hsl(var(--bc) / 0.2);
  border-radius: 10px;
}
.bg-dot-pattern {
  background-image: radial-gradient(hsl(var(--bc) / 0.05) 1px, transparent 1px);
  background-size: 20px 20px;
}
</style>

<script setup>
const route = useRoute()

const question = ref('')
const answer = ref('')
const loading = ref(false)
let eventSource = null

onMounted(() => {
  if (route.query.question) {
    question.value = route.query.question
    startChat()
  }
})

const startChat = () => {
  if (!question.value.trim() || loading.value) return

  answer.value = ''
  loading.value = true

  const q = encodeURIComponent(question.value)

  eventSource = new EventSource(`http://localhost:2121/chat?question=${q}`)

  eventSource.onmessage = (event) => {
    answer.value += event.data
  }

  eventSource.addEventListener('done', () => {
    loading.value = false
    eventSource.close()
  })
}
</script>
