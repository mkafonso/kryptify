import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Button, Input } from "../component";
import { createAccountSchema } from "../util";

type CreateAccountFormData = Required<z.infer<typeof createAccountSchema>>;

export function Login() {
  const {
    register,
    handleSubmit,
    getValues,
    formState: { errors },
  } = useForm<CreateAccountFormData>({
    resolver: zodResolver(createAccountSchema),
  });

  const submitForm = async (data: CreateAccountFormData) => {
    console.log(data);
  };

  return (
    <form
      autoComplete="off"
      autoCorrect="off"
      autoCapitalize="off"
      className="max-w-sm w-full p-4"
      onSubmit={handleSubmit(submitForm)}
    >
      <h1 className="text-2xl mb-1 font-bold">Sign in</h1>
      <p className="mb-6 text-base">Welcome back!</p>

      <div className="flex flex-col gap-4">
        <Input
          name="email"
          label="Email address"
          placeholder="Enter your email"
          type="email"
          hasValue={!!getValues("email")}
          register={register}
          error={errors.email}
        />

        <Input
          name="password"
          type="password"
          label="Password"
          placeholder="Password"
          hasValue={!!getValues("password")}
          register={register}
          error={errors.password}
        />

        <a
          href="#/forgot-password"
          className="w-max text-sm text-primary-500 transition"
        >
          Forgot password?
        </a>

        <Button type="submit">Sign in</Button>

        <p className="flex items-center text-sm">
          Don't have an account?
          <a
            href="#/login"
            className="inline-block ml-1 text-sm font-bold text-primary-500 transition"
          >
            Sign up
          </a>
        </p>
      </div>
    </form>
  );
}
