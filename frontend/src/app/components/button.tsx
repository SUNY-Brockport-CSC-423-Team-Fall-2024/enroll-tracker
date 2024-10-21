import styles from "./styles.module.css";

interface ButtonWithCallbackProps {
  btnTitle: string;
  onClick: (...args: any[]) => any;
}

const Button: React.FC<ButtonWithCallbackProps> = ({ onClick, btnTitle }) => {
  return (
    <div className={styles.button} onClick={onClick}>
      <p>{btnTitle}</p>
    </div>
  );
};

export default Button;
