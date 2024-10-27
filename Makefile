.PHONY: create-vue install setup run create-structure create-files

# デフォルト設定
PROJECT_NAME ?= my-vue-app
SRC_DIR = $(PROJECT_NAME)/src

create-vue:
	npm create vue@3 $(PROJECT_NAME) -- \
		--typescript \
		--jsx \
		--router \
		--pinia \
		--vitest \
		--eslint \
		--prettier \
		--force

install:
	cd $(PROJECT_NAME) && npm install

create-structure:
	mkdir -p $(SRC_DIR)/assets
	mkdir -p $(SRC_DIR)/components/common
	mkdir -p $(SRC_DIR)/components/home
	mkdir -p $(SRC_DIR)/components/users
	mkdir -p $(SRC_DIR)/components/auth
	mkdir -p $(SRC_DIR)/layouts
	mkdir -p $(SRC_DIR)/router
	mkdir -p $(SRC_DIR)/stores
	mkdir -p $(SRC_DIR)/types
	mkdir -p $(SRC_DIR)/views/users
	mkdir -p $(SRC_DIR)/views/auth
	mkdir -p $(SRC_DIR)/composables

create-files: create-layouts create-router create-views create-components create-stores

create-layouts:
	# DefaultLayout
	echo "<template>\n\
  <div class=\"min-h-screen bg-gray-50\">\n\
    <Navigation />\n\
    <main>\n\
      <slot></slot>\n\
    </main>\n\
  </div>\n\
</template>\n\
\n\
<script setup lang=\"ts\">\n\
import Navigation from '@/components/common/Navigation.vue';\n\
</script>" > $(SRC_DIR)/layouts/DefaultLayout.vue

create-router:
	echo "import { createRouter, createWebHistory } from 'vue-router';\n\
import type { RouteRecordRaw } from 'vue-router';\n\
\n\
const routes: RouteRecordRaw[] = [\n\
  {\n\
    path: '/',\n\
    name: 'home',\n\
    component: () => import('@/views/HomeView.vue'),\n\
  },\n\
  {\n\
    path: '/users',\n\
    name: 'users',\n\
    component: () => import('@/views/users/UsersLayout.vue'),\n\
    children: [\n\
      {\n\
        path: '',\n\
        name: 'users-list',\n\
        component: () => import('@/views/users/UsersList.vue'),\n\
      },\n\
      {\n\
        path: ':id',\n\
        name: 'user-detail',\n\
        component: () => import('@/views/users/UserDetail.vue'),\n\
        props: true,\n\
      },\n\
    ],\n\
  },\n\
  {\n\
    path: '/auth',\n\
    name: 'auth',\n\
    component: () => import('@/views/auth/AuthLayout.vue'),\n\
    children: [\n\
      {\n\
        path: 'login',\n\
        name: 'login',\n\
        component: () => import('@/views/auth/LoginView.vue'),\n\
      },\n\
      {\n\
        path: 'register',\n\
        name: 'register',\n\
        component: () => import('@/views/auth/RegisterView.vue'),\n\
      },\n\
    ],\n\
  },\n\
];\n\
\n\
const router = createRouter({\n\
  history: createWebHistory(import.meta.env.BASE_URL),\n\
  routes,\n\
});\n\
\n\
export default router;" > $(SRC_DIR)/router/index.ts

create-views:
	# Home View
	echo "<template>\n\
  <DefaultLayout>\n\
    <div class=\"container mx-auto px-4 py-8\">\n\
      <h1 class=\"text-2xl font-bold mb-4\">Home</h1>\n\
    </div>\n\
  </DefaultLayout>\n\
</template>\n\
\n\
<script setup lang=\"ts\">\n\
import DefaultLayout from '@/layouts/DefaultLayout.vue';\n\
</script>" > $(SRC_DIR)/views/HomeView.vue

	# Users Layout
	echo "<template>\n\
  <DefaultLayout>\n\
    <div class=\"container mx-auto px-4 py-8\">\n\
      <router-view />\n\
    </div>\n\
  </DefaultLayout>\n\
</template>\n\
\n\
<script setup lang=\"ts\">\n\
import DefaultLayout from '@/layouts/DefaultLayout.vue';\n\
</script>" > $(SRC_DIR)/views/users/UsersLayout.vue

	# Users List
	echo "<template>\n\
  <div>\n\
    <h1 class=\"text-2xl font-bold mb-4\">Users List</h1>\n\
    <div class=\"grid gap-4\">\n\
      <div v-for=\"id in 5\" :key=\"id\" class=\"p-4 bg-white rounded shadow\">\n\
        <router-link :to=\"{ name: 'user-detail', params: { id } }\" class=\"text-blue-600 hover:underline\">\n\
          User {{ id }}\n\
        </router-link>\n\
      </div>\n\
    </div>\n\
  </div>\n\
</template>" > $(SRC_DIR)/views/users/UsersList.vue

	# User Detail
	echo "<template>\n\
  <div>\n\
    <h1 class=\"text-2xl font-bold mb-4\">User Details</h1>\n\
    <div class=\"bg-white rounded shadow p-4\">\n\
      <p>User ID: {{ id }}</p>\n\
    </div>\n\
  </div>\n\
</template>\n\
\n\
<script setup lang=\"ts\">\n\
defineProps<{ id: string }>();\n\
</script>" > $(SRC_DIR)/views/users/UserDetail.vue

	# Auth Layout
	echo "<template>\n\
  <div class=\"min-h-screen bg-gray-50 flex justify-center items-center\">\n\
    <router-view />\n\
  </div>\n\
</template>" > $(SRC_DIR)/views/auth/AuthLayout.vue

	# Login View
	echo "<template>\n\
  <div class=\"max-w-md w-full mx-auto bg-white p-8 rounded-lg shadow-md\">\n\
    <h2 class=\"text-2xl font-bold mb-6\">Login</h2>\n\
    <form @submit.prevent=\"handleSubmit\">\n\
      <div class=\"space-y-4\">\n\
        <div>\n\
          <label class=\"block text-sm font-medium text-gray-700\">Email</label>\n\
          <input type=\"email\" v-model=\"email\" class=\"mt-1 block w-full rounded-md border-gray-300 shadow-sm\">\n\
        </div>\n\
        <div>\n\
          <label class=\"block text-sm font-medium text-gray-700\">Password</label>\n\
          <input type=\"password\" v-model=\"password\" class=\"mt-1 block w-full rounded-md border-gray-300 shadow-sm\">\n\
        </div>\n\
        <button type=\"submit\" class=\"w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700\">Login</button>\n\
      </div>\n\
    </form>\n\
  </div>\n\
</template>\n\
\n\
<script setup lang=\"ts\">\n\
import { ref } from 'vue';\n\
const email = ref('');\n\
const password = ref('');\n\
\n\
const handleSubmit = () => {\n\
  console.log({ email: email.value, password: password.value });\n\
};\n\
</script>" > $(SRC_DIR)/views/auth/LoginView.vue

	# Register View
	echo "<template>\n\
  <div class=\"max-w-md w-full mx-auto bg-white p-8 rounded-lg shadow-md\">\n\
    <h2 class=\"text-2xl font-bold mb-6\">Register</h2>\n\
    <form @submit.prevent=\"handleSubmit\">\n\
      <div class=\"space-y-4\">\n\
        <div>\n\
          <label class=\"block text-sm font-medium text-gray-700\">Name</label>\n\
          <input type=\"text\" v-model=\"name\" class=\"mt-1 block w-full rounded-md border-gray-300 shadow-sm\">\n\
        </div>\n\
        <div>\n\
          <label class=\"block text-sm font-medium text-gray-700\">Email</label>\n\
          <input type=\"email\" v-model=\"email\" class=\"mt-1 block w-full rounded-md border-gray-300 shadow-sm\">\n\
        </div>\n\
        <div>\n\
          <label class=\"block text-sm font-medium text-gray-700\">Password</label>\n\
          <input type=\"password\" v-model=\"password\" class=\"mt-1 block w-full rounded-md border-gray-300 shadow-sm\">\n\
        </div>\n\
        <button type=\"submit\" class=\"w-full bg-blue-600 text-white py-2 px-4 rounded-md hover:bg-blue-700\">Register</button>\n\
      </div>\n\
    </form>\n\
  </div>\n\
</template>\n\
\n\
<script setup lang=\"ts\">\n\
import { ref } from 'vue';\n\
const name = ref('');\n\
const email = ref('');\n\
const password = ref('');\n\
\n\
const handleSubmit = () => {\n\
  console.log({ name: name.value, email: email.value, password: password.value });\n\
};\n\
</script>" > $(SRC_DIR)/views/auth/RegisterView.vue

create-components:
	# Navigation
	echo "<template>\n\
  <nav class=\"bg-white shadow\">\n\
    <div class=\"container mx-auto px-4\">\n\
      <div class=\"flex justify-between h-16\">\n\
        <div class=\"flex space-x-4 items-center\">\n\
          <router-link to=\"/\" class=\"text-gray-700 hover:text-gray-900\">Home</router-link>\n\
          <router-link to=\"/users\" class=\"text-gray-700 hover:text-gray-900\">Users</router-link>\n\
        </div>\n\
        <div class=\"flex items-center space-x-4\">\n\
          <router-link to=\"/auth/login\" class=\"text-gray-700 hover:text-gray-900\">Login</router-link>\n\
          <router-link to=\"/auth/register\" class=\"text-gray-700 hover:text-gray-900\">Register</router-link>\n\
        </div>\n\
      </div>\n\
    </div>\n\
  </nav>\n\
</template>" > $(SRC_DIR)/components/common/Navigation.vue

create-stores:
	echo "import { defineStore } from 'pinia';\n\
\n\
interface User {\n\
  id: number;\n\
  email: string;\n\
  name: string;\n\
}\n\
\n\
export const useAuthStore = defineStore('auth', {\n\
  state: () => ({\n\
    user: null as User | null,\n\
    loading: false,\n\
  }),\n\
\n\
  getters: {\n\
    isAuthenticated: (state) => !!state.user,\n\
  },\n\
\n\
  actions: {\n\
    async login(email: string, password: string) {\n\
      this.loading = true;\n\
      try {\n\
        // API call would go here\n\
        this.user = {\n\
          id: 1,\n\
          email,\n\
          name: 'Test User',\n\
        };\n\
      } finally {\n\
        this.loading = false;\n\
      }\n\
    },\n\
\n\
    logout() {\n\
      this.user = null;\n\
    },\n\
  },\n\
});" > $(SRC_DIR)/stores/auth.ts

run:
	cd $(PROJECT_NAME) && npm run dev

setup: create-vue install create-structure create-files run