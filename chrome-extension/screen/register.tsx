import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Button, Input } from "../component";
import { createAccountSchema } from "../util";

type CreateAccountFormData = Required<z.infer<typeof createAccountSchema>>;

export function Register() {
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
      <h1 className="text-2xl mb-1 font-bold">Register</h1>
      <p className="mb-6 text-base">Create your account</p>

      <div className="flex flex-col gap-4">
        <Input
          name="name"
          label="Name"
          placeholder="Enter your name"
          hasValue={!!getValues("name")}
          register={register}
          error={errors.name}
        />

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

        <Button type="submit">Register</Button>

        <p className="flex items-center text-sm">
          Already have an account?
          <a
            href="#/login"
            className="inline-block ml-1 text-sm font-bold text-primary-500 transition"
          >
            Sign in
          </a>
        </p>

        <p className="text-sm">
          By clicking on create account, you agree to the{" "}
          <a
            href="#"
            target="_blank"
            className="text-primary-500 font-semibold"
          >
            terms of use
          </a>{" "}
          and{" "}
          <a
            href="#"
            target="_blank"
            className="text-primary-500 font-semibold"
          >
            privacy policies
          </a>
          .
        </p>
      </div>
    </form>
  );
}
