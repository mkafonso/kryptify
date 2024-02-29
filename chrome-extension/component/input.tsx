import { UseFormRegister } from "react-hook-form";

interface InputFormProps extends React.InputHTMLAttributes<HTMLInputElement> {
  label: string;
  name: string;
  error?: any;
  hasValue?: boolean;
  updating?: boolean;
  register: UseFormRegister<any>;
}

export function Input(props: InputFormProps) {
  const { label, hasValue, name, error, register, updating, ...rest } = props;

  return (
    <div className="block">
      {!updating && (
        <label className="block text-base font-medium mb-1" htmlFor={name}>
          {label}
        </label>
      )}

      <input
        id={name}
        data-error={!!error}
        data-isfilled={!!hasValue && !error}
        className="text-main-800 text-base bg-white/5 border border-main-500/20 px-4 h-11 w-full rounded-md shadow-sm placeholder-main-500
        data-[error=true]:border-2 data-[error=true]:border-danger-500 data-[error=true]:focus:outline-none
        data-[isfilled=true]:border-2 data-[isfilled=true]:!border-success-500 data-[isfilled=true]:focus:outline-none"
        {...rest}
        {...register(name)}
      />

      {!!error && (
        <span className="text-danger-500 text-sm text-semibold block">
          {error.message}
        </span>
      )}
    </div>
  );
}
