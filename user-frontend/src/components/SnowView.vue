<script setup lang="ts">
import { onMounted, onBeforeUnmount } from 'vue'

let canvas: HTMLCanvasElement
let ctx: CanvasRenderingContext2D
let animationId: number

const flakes: any[] = []
const COUNT = 120

function resize() {
  canvas.width = window.innerWidth
  canvas.height = window.innerHeight
}

function createFlakes() {
  flakes.length = 0
  for (let i = 0; i < COUNT; i++) {
    flakes.push({
      x: Math.random() * canvas.width,
      y: Math.random() * canvas.height,
      r: Math.random() * 3 + 1,
      d: Math.random() + 0.5,
      o: Math.random() * 0.5 + 0.3,
    })
  }
}

function draw() {
  ctx.clearRect(0, 0, canvas.width, canvas.height)
  ctx.fillStyle = 'rgba(255,255,255,0.8)'
  ctx.beginPath()

  flakes.forEach((f) => {
    ctx.globalAlpha = f.o
    ctx.moveTo(f.x, f.y)
    ctx.arc(f.x, f.y, f.r, 0, Math.PI * 2)
  })

  ctx.fill()
  update()
  animationId = requestAnimationFrame(draw)
}

function update() {
  flakes.forEach((f) => {
    f.y += f.d
    if (f.y > canvas.height) {
      f.y = 0
      f.x = Math.random() * canvas.width
    }
  })
}

onMounted(() => {
  canvas = document.getElementById('snow') as HTMLCanvasElement
  ctx = canvas.getContext('2d')!
  resize()
  createFlakes()
  draw()
  window.addEventListener('resize', resize)
})

onBeforeUnmount(() => {
  cancelAnimationFrame(animationId)
  window.removeEventListener('resize', resize)
})
</script>

<template>
  <canvas id="snow" class="fixed inset-0 pointer-events-none z-50" />
</template>
