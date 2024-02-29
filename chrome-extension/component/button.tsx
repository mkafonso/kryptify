interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  children: React.ReactNode;
}

export function Button(props: ButtonProps) {
  const { children, ...rest } = props;

  return (
    <button
      className="text-main-100 text-base flex gap-4 items-center justify-center rounded-md radius-md h-11 hover:bg-primary-400 bg-primary-500 font-semibold transition"
      {...rest}
    >
      {children}
    </button>
  );
}
