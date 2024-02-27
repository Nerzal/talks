public interface IAnimal {
	void animalSound();
}

public class Theelinger : IAnimal {
	public void animalSound() {
		Console.WriteLine("Tobigeräusche")
	}
}