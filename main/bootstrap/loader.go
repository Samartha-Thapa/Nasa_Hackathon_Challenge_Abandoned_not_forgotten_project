components {
  id: "loader"
  component: "/main/bootstrap/loader.script"
}
embedded_components {
  id: "proxy_earth"
  type: "collectionproxy"
  data: "collection: \"/main/main.collection\"\n"
  ""
}
embedded_components {
  id: "proxy_moon"
  type: "collectionproxy"
  data: "collection: \"/main/moon.collection\"\n"
  ""
}
embedded_components {
  id: "proxy_mars"
  type: "collectionproxy"
  data: "collection: \"/main/moon.collection\"\n"
  ""
}
embedded_components {
  id: "proxy_opening"
  type: "collectionproxy"
  data: "collection: \"/main/earth_opening_sequence.collection\"\n"
  ""
}
