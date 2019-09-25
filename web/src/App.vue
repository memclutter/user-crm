<template>
  <v-app>
    <v-app-bar dark color="accent">
      <v-app-bar-nav-icon></v-app-bar-nav-icon>

      <v-toolbar-title>User CRM</v-toolbar-title>

      <div class="flex-grow-1"></div>

      <v-toolbar-items>
        <v-btn text :to="{name: 'users'}">
          <v-icon>fa-users</v-icon>
        </v-btn>

        <v-btn text :to="{name: 'countries'}">
          <v-icon>fa-globe-americas</v-icon>
        </v-btn>

        <v-btn text href="https://github.com/memclutter/user-crm" target="_blank">
          <v-icon>fab fa-github</v-icon>
        </v-btn>
      </v-toolbar-items>
    </v-app-bar>
    <v-content>
      <router-view></router-view>
    </v-content>
  </v-app>
</template>