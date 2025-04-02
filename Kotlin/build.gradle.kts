plugins {
    kotlin("jvm") version "2.1.10"
    kotlin("plugin.serialization") version "2.1.20"
    application
}

group = "org.example"
version = "1.0-SNAPSHOT"

val ktor_version: String by project
val kdor_version: String by project

application {
    mainClass = "MainKt" 
}

repositories {
    mavenCentral()
    // maven("https://repo.kord.dev/snapshots")
}

dependencies {
    implementation("io.ktor:ktor-client-core:$ktor_version")
    implementation("io.ktor:ktor-client-cio:$ktor_version")
    implementation("io.ktor:ktor-client-content-negotiation:$ktor_version")
    implementation("io.ktor:ktor-serialization-kotlinx-json:$ktor_version")
    implementation("org.jetbrains.kotlinx:kotlinx-serialization-json:1.8.1")
    implementation("com.typesafe:config:1.4.3")

    // implementation("dev.kord:kord-core:$kdor_version")
}

kotlin {
    jvmToolchain(21)
}