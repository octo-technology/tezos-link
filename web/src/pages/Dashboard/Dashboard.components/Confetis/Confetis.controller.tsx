let loopId: number, canvas: any, ctx: any, unit: number, confetis: Confeti[]

const CSS_COLOR_NAMES = [
  'AliceBlue',
  'AntiqueWhite',
  'Aqua',
  'Aquamarine',
  'Azure',
  'Beige',
  'Bisque',
  'Black',
  'BlanchedAlmond',
  'Blue',
  'BlueViolet',
  'Brown',
  'BurlyWood',
  'CadetBlue',
  'Chartreuse',
  'Chocolate',
  'Coral',
  'CornflowerBlue',
  'Cornsilk',
  'Crimson',
  'Cyan',
  'DarkBlue',
  'DarkCyan',
  'DarkGoldenRod',
  'DarkGray',
  'DarkGrey',
  'DarkGreen',
  'DarkKhaki',
  'DarkMagenta',
  'DarkOliveGreen',
  'Darkorange',
  'DarkOrchid',
  'DarkRed',
  'DarkSalmon',
  'DarkSeaGreen',
  'DarkSlateBlue',
  'DarkSlateGray',
  'DarkSlateGrey',
  'DarkTurquoise',
  'DarkViolet',
  'DeepPink',
  'DeepSkyBlue',
  'DimGray',
  'DimGrey',
  'DodgerBlue',
  'FireBrick',
  'FloralWhite',
  'ForestGreen',
  'Fuchsia',
  'Gainsboro',
  'GhostWhite',
  'Gold',
  'GoldenRod',
  'Gray',
  'Grey',
  'Green',
  'GreenYellow',
  'HoneyDew',
  'HotPink',
  'IndianRed',
  'Indigo',
  'Ivory',
  'Khaki',
  'Lavender',
  'LavenderBlush',
  'LawnGreen',
  'LemonChiffon',
  'LightBlue',
  'LightCoral',
  'LightCyan',
  'LightGoldenRodYellow',
  'LightGray',
  'LightGrey',
  'LightGreen',
  'LightPink',
  'LightSalmon',
  'LightSeaGreen',
  'LightSkyBlue',
  'LightSlateGray',
  'LightSlateGrey',
  'LightSteelBlue',
  'LightYellow',
  'Lime',
  'LimeGreen',
  'Linen',
  'Magenta',
  'Maroon',
  'MediumAquaMarine',
  'MediumBlue',
  'MediumOrchid',
  'MediumPurple',
  'MediumSeaGreen',
  'MediumSlateBlue',
  'MediumSpringGreen',
  'MediumTurquoise',
  'MediumVioletRed',
  'MidnightBlue',
  'MintCream',
  'MistyRose',
  'Moccasin',
  'NavajoWhite',
  'Navy',
  'OldLace',
  'Olive',
  'OliveDrab',
  'Orange',
  'OrangeRed',
  'Orchid',
  'PaleGoldenRod',
  'PaleGreen',
  'PaleTurquoise',
  'PaleVioletRed',
  'PapayaWhip',
  'PeachPuff',
  'Peru',
  'Pink',
  'Plum',
  'PowderBlue',
  'Purple',
  'Red',
  'RosyBrown',
  'RoyalBlue',
  'SaddleBrown',
  'Salmon',
  'SandyBrown',
  'SeaGreen',
  'SeaShell',
  'Sienna',
  'Silver',
  'SkyBlue',
  'SlateBlue',
  'SlateGray',
  'SlateGrey',
  'Snow',
  'SpringGreen',
  'SteelBlue',
  'Tan',
  'Teal',
  'Thistle',
  'Tomato',
  'Turquoise',
  'Violet',
  'Wheat',
  'White',
  'WhiteSmoke',
  'Yellow',
  'YellowGreen'
]

// Customization
const amount = 100
const gravity = 0.098

class Confeti {
  opacity: number
  speedOpacity: number
  heigth: number
  heigthv: number
  rotation: number
  rotationv: number
  x: number
  y: number
  vx: number
  vy: number
  width: number
  color: string

  constructor() {
    this.opacity = 0
    this.speedOpacity = 0
    this.heigth = 0
    this.heigthv = 0
    this.rotation = 0
    this.rotationv = 0
    this.x = 0
    this.y = 0
    this.vx = 0
    this.vy = 0
    this.width = 0
    this.color = ''

    this.randomInitialise()
  }

  update = () => {
    if (this.opacity < 0) {
      this.randomInitialise()
    } else {
      this.opacity -= this.speedOpacity
    }
    this.heigthv = this.heigth < this.width && this.heigth > 0 ? this.heigthv : -this.heigthv
    this.heigth += this.heigthv
    this.rotation += this.rotationv
    this.x = this.x + this.vx
    this.y = this.y + this.vy
    this.vy += gravity
  }

  draw = () => {
    ctx.save()

    ctx.beginPath()
    ctx.globalAlpha = this.opacity
    ctx.fillStyle = this.color
    ctx.translate(this.x, this.y)
    ctx.rotate((this.rotation * Math.PI) / 180)
    ctx.ellipse(0 - this.width / 2, 0 - this.heigth / 2, Math.abs(this.width), Math.abs(this.heigth), 0, 0, 2 * Math.PI)
    ctx.fill()
    ctx.closePath()

    ctx.restore()
  }

  randomInitialise = () => {
    this.x = canvas.width / 2
    this.y = canvas.height / 2
    this.vx = -2 + Math.random() * 4
    this.vy = Math.random() * -5
    this.width = (Math.random() * unit) / 20
    this.heigth = Math.random() * this.width
    this.heigthv = Math.random() * 3
    this.rotation = Math.random() * 360
    this.rotationv = Math.random() * 2
    this.color = CSS_COLOR_NAMES[Math.floor(Math.random() * CSS_COLOR_NAMES.length)]
    this.opacity = 1
    this.speedOpacity = Math.random() * 0.001
  }
}

function calculateUnit() {
  const maxLength = Math.min(canvas.width, canvas.height)
  return maxLength / 2.5
}

function initialiseElements() {
  confetis = []
  confetis.push(new Confeti())
}

function startAnimation() {
  // eslint-disable-next-line @typescript-eslint/no-this-alias
  loopId = requestAnimationFrame(animationLoop)
}

export function showConfetis() {
  // Get canvas & context
  canvas = document.getElementById('confetis')
  // @ts-ignore
  if (canvas) {
    canvas.style.zIndex = '999'
    canvas.width = window.innerWidth
    canvas.height = window.innerHeight
    ctx = canvas.getContext('2d')
    // initialise animation & variables
    initialiseElements()
    startAnimation()
  }
}

function drawBackground() {
  if (loopId >= 30) {
    confetis = confetis.slice(0, confetis.length - 2)
  }

  if (loopId < 30) {
    if (confetis.length < amount - 1) {
      for (let i = 0; i <= 10; i++) {
        confetis.push(new Confeti())
      }
    }
  }

  if (loopId < 150) {
    for (let i = 0; i <= confetis.length - 1; i++) {
      confetis[i].draw()
      confetis[i].update()
    }
  }
}

function drawScene() {
  drawBackground()
}

function animationLoop() {
  if (loopId >= 150) {
    // @ts-ignore
    canvas.style.zIndex = '-1'
    return cancelAnimationFrame(loopId)
  }
  // 1 - Clear & resize
  ctx.clearRect(0, 0, canvas.width, canvas.height)
  canvas.width = 400
  canvas.height = 200
  ctx = canvas.getContext('2d')
  unit = calculateUnit()
  // 2 Draw & Move
  drawScene()
  // call again mainloop after 16.6ms
  // (corresponds to 60 frames/second)
  loopId = requestAnimationFrame(animationLoop)
}
