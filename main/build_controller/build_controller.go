components {
  id: "build_controller"
  component: "/main/build_controller/build_controller.script"
}
embedded_components {
  id: "factory_solar"
  type: "factory"
  data: "prototype: \"/main/buildings/building_solar.go\"\n"
  ""
}
embedded_components {
  id: "factory_museum"
  type: "factory"
  data: "prototype: \"/main/buildings/building_museum.go\"\n"
  ""
}
embedded_components {
  id: "factory_research"
  type: "factory"
  data: "prototype: \"/main/buildings/building_research.go\"\n"
  ""
}
embedded_components {
  id: "factory_rocket"
  type: "factory"
  data: "prototype: \"/main/buildings/building_rocket.go\"\n"
  ""
}
