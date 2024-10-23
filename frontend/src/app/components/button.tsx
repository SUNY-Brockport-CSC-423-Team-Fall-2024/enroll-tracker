interface ButtonWithCallbackProps {
  btnTitle: string;
  onClick: (...args: any[]) => any;
}

const Button: React.FC<ButtonWithCallbackProps> = ({ onClick, btnTitle }) => {
  return (
    <div onClick={onClick}>
      <p>{btnTitle}</p>
    </div>
  );
};

export default Button;
