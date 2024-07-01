components {
  id: "cell"
  component: "/assets/scripts/cell.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "blue"
  type: "sprite"
  data: "default_animation: \"blue_cell\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlasses/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: -0.4
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 0.75
    y: 0.75
    z: 1.0
  }
}
embedded_components {
  id: "green"
  type: "sprite"
  data: "default_animation: \"green_cell\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlasses/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: -0.4
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 0.75
    y: 0.75
    z: 1.0
  }
}
embedded_components {
  id: "red"
  type: "sprite"
  data: "default_animation: \"red_cell\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlasses/sprites.atlas\"\n"
  "}\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: -0.4
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  scale {
    x: 0.75
    y: 0.75
    z: 1.0
  }
}
embedded_components {
  id: "left"
  type: "factory"
  data: "prototype: \"/assets/objects/left_arrow.go\"\n"
  "load_dynamically: false\n"
  "dynamic_prototype: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "down"
  type: "factory"
  data: "prototype: \"/assets/objects/down_arrow.go\"\n"
  "load_dynamically: false\n"
  "dynamic_prototype: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "right"
  type: "factory"
  data: "prototype: \"/assets/objects/right_arrow.go\"\n"
  "load_dynamically: false\n"
  "dynamic_prototype: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "up"
  type: "factory"
  data: "prototype: \"/assets/objects/up_arrow.go\"\n"
  "load_dynamically: false\n"
  "dynamic_prototype: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
