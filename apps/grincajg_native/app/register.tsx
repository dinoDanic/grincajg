import { Spacing } from "@/constants"
import { graphql } from "@/generated"
import { CreateUserDocument, CreateUserInput } from "@/generated/graphql"
import { _client } from "@/lib"
import { zodResolver } from "@hookform/resolvers/zod"
import { useMutation, useQuery } from "@tanstack/react-query"
import { router } from "expo-router"
import { Controller, useForm } from "react-hook-form"
import { Button, Text, View } from "react-native"
import { z } from "zod"

import { Input } from "@/components/ui/input"

const resolver = z
  .object({
    email: z.string().email(),
    password: z.string().min(1),
    passwordConfirm: z.string().min(1),
    name: z.string().min(1),
  })
  .refine((data) => data.password === data.passwordConfirm, {
    message: "Passwords do not match",
    path: ["passwordConfirm"],
  })

const loginFormResolver = zodResolver(resolver)

const RegisterMutationDucoment = graphql(/* GraphQL */ `
  mutation createUser($input: CreateUserInput!) {
    createUser(input: $input) {
      name
      id
    }
  }
`)

const meDocument = graphql(/* GraphQL */ `
  query me {
    me {
      name
    }
  }
`)

export default function App() {
  const {
    control,
    handleSubmit,
    formState: { errors },
    setError,
  } = useForm<CreateUserInput>({
    resolver: loginFormResolver,
    defaultValues: {
      email: "",
      password: "",
      passwordConfirm: "",
      name: "",
    },
  })

  // const registerMutataion = useMutation( )
  const { data } = useQuery(["me"], async () => _client.request(meDocument))

  const onSubmit = async (data: CreateUserInput) => {
    try {
      const res = await _client.request(CreateUserDocument, {
        input: data,
      })
      if (res.createUser.id) {
        router.replace("/sign-in")
      }
    } catch (err) {
      console.log(err)

      setError("name", { message: "Something went wrong" })
    }
  }

  console.log(errors)

  return (
    <View style={{ flex: 1, justifyContent: "center", alignItems: "center", gap: Spacing.md, padding: Spacing.md }}>
      <Controller
        control={control}
        rules={{
          required: true,
        }}
        render={({ field: { onChange, onBlur, value } }) => (
          <Input placeholder="email" onBlur={onBlur} onChangeText={onChange} value={value} />
        )}
        name="email"
      />
      {errors.email && <Text>{errors.email.message}</Text>}

      <Controller
        control={control}
        rules={{
          maxLength: 100,
        }}
        render={({ field: { onChange, onBlur, value } }) => (
          <Input placeholder="Password" onBlur={onBlur} onChangeText={onChange} value={value} secureTextEntry />
        )}
        name="password"
      />
      {errors.password && <Text>{errors.password.message}</Text>}

      <Controller
        control={control}
        rules={{
          maxLength: 100,
        }}
        render={({ field: { onChange, onBlur, value } }) => (
          <Input placeholder="Repeat password" onBlur={onBlur} onChangeText={onChange} value={value} secureTextEntry />
        )}
        name="passwordConfirm"
      />
      {errors.passwordConfirm && <Text>{errors.passwordConfirm.message}</Text>}
      <Controller
        control={control}
        rules={{
          maxLength: 100,
        }}
        render={({ field: { onChange, onBlur, value } }) => (
          <Input placeholder="Name" onBlur={onBlur} onChangeText={onChange} value={value} />
        )}
        name="name"
      />
      {errors.name && <Text>{errors.name.message}</Text>}

      <Button title="register" onPress={handleSubmit(onSubmit)} />
    </View>
  )
}
