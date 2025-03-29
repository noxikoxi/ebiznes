plugins {
    kotlin("jvm") version "2.1.10"
}

group = "org.example"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

val ktor_version: String by project
val logback_version: String by project

dependencies {
    testImplementation(kotlin("test"))

    // klient ktor
    implementation("io.ktor:ktor-client-core:$ktor_version")
    implementation("io.ktor:ktor-client-cio:$ktor_version")

    // logger
    implementation("ch.qos.logback:logback-classic:$logback_version")

    // for discord api
    implementation("io.ktor:ktor-client-websockets:2.3.8")
    implementation("io.ktor:ktor-client-serialization:2.3.8")
}

tasks.test {
    useJUnitPlatform()
}
kotlin {
    jvmToolchain(21)
}