plugins {
    kotlin("jvm") version "2.1.10"
    application
}

group = "org.example"
version = "1.0-SNAPSHOT"

application {
    mainClass = "MainKt" 
}

repositories {
    mavenCentral()
}

// dependencies {
    //
// }

kotlin {
    jvmToolchain(21)
}