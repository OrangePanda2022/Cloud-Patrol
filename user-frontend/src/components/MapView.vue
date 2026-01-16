<template>
  <div class="relative w-full">
    <!-- 状态显示 -->
    <div
      class="absolute top-6 left-6 z-1 bg-green-900 text-white text-md px-3 py-1 rounded-md flex items-center gap-2"
    >
      正在飞行
    </div>

    <div class="card p-4"><div id="map" class="w-full h-[400px] rounded-md shadow"></div></div>
  </div>
</template>

<script setup>
let map = null
let marker = null
let polyline = null
let passedPolyline = null

// 示例飞行路径（经纬度）
const path = [
  [114.3311545317969, 30.57486180366169],
  [114.32922154983505, 30.57666879262041],
  [114.33347877702505, 30.58011820722278],
  [114.33498110331365, 30.578640521692186],
  [114.3328775006871, 30.57671578345881],
]

const loadAMap = () => {
  return new Promise((resolve) => {
    if (window.AMap) return resolve()
    const script = document.createElement('script')
    script.src =
      'https://webapi.amap.com/maps?v=2.0&key=26b0104caf88469b96f4a2673b85d2fc&plugin=AMap.MoveAnimation'
    script.onload = resolve
    document.head.appendChild(script)
  })
}

onMounted(async () => {
  await loadAMap()

  map = new AMap.Map('map', {
    zoom: 18,
    center: path[0],
  })

  polyline = new AMap.Polyline({
    map,
    path,
    strokeColor: '#00bcd4',
    strokeWeight: 4,
  })

  passedPolyline = new AMap.Polyline({
    map,
    strokeColor: '#ff5722',
    strokeWeight: 4,
  })

  marker = new AMap.Marker({
    map,
    position: path[0],
    icon: new AMap.Icon({
      image: new URL('@/assets/drone.png', import.meta.url).href,
      size: new AMap.Size(50, 50),
      imageSize: new AMap.Size(50, 50),
    }),
    offset: new AMap.Pixel(-25, -25),
  })

  marker.on('moving', (e) => {
    passedPolyline.setPath(e.passedPath)
    map.setCenter(e.target.getPosition(), false)
    map.panTo(e.target.getPosition(), 300)
  })
})

const start = () => {
  if (!marker) return
  passedPolyline.setPath([])
  marker.moveAlong(path, {
    duration: 50000,
    autoRotation: true,
  })
}

const reset = () => {
  if (!marker) return
  marker.stopMove()
  marker.setPosition(path[0])
  passedPolyline.setPath([])
}
onMounted(() => {
  setTimeout(() => {
    start()
  }, 1000)
})
</script>

<style scoped>
#map {
  min-height: 400px;
}
</style>
