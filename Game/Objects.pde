
class Planet {
  float x, y;
  int owner;
  String name;
  float[] resources;

  Planet(float x, float y, int owner, String name, float[] resources) {
    this.x = x;
    this.y = y;
    this.owner = owner;
    this.name = name;
    this.resources = resources;
  }
}

class Ship {
  String name;
  int speed;
  int owner;
  float x, y;
  float destX, destY;

  Ship(String name, int speed, int owner, float x, float y, float destX, float destY) {
    this.name = name;
    this.speed = speed;
    this.owner = owner;
    this.x = x;
    this.y = y;
    this.destX = destX;
    this.destY = destY;
  }
}
