# Setting Up SSH Authentication for GitHub

SSH authentication is a secure method to connect to GitHub without needing to enter your password or personal access token each time. Instead, you create an SSH key pair (public and private keys) and add the public key to your GitHub account. Here’s how to set it up:

## Step 1: Check for Existing SSH Keys
First, check if you already have an SSH key pair on your machine:
```bash
ls -al ~/.ssh
```
If you see files like id_rsa and id_rsa.pub, you already have SSH keys and can skip to **Step 3**. If not, proceed to **Step 2** to create them.

## Step 2: Generate a New SSH Key Pair
1. Run the following command to generate a new SSH key. Replace <your-email@example.com> with your GitHub email:
```bash
ssh-keygen -t rsa -b 4096 -C "<your-email@example.com>"
```
2. When prompted to "Enter a file in which to save the key," press Enter to save it to the default location (~/.ssh/id_rsa).
3. You can add a passphrase for additional security, but it’s optional.

## Step 3: Start the SSH Agent and Add Your Key
1. Start the SSH agent in the background:
```bash
eval "$(ssh-agent -s)"
```
2. Add your SSH private key to the SSH agent:
```bash
ssh-add ~/.ssh/id_rsa
```

## Step 4: Add Your SSH Public Key to GitHub
1. Copy the contents of your public key to your clipboard:
```bash
cat ~/.ssh/id_rsa.pub
```
This will display the key; select and copy it.

2. Go to your GitHub account:
    - Navigate to **Settings** > **SSH and GPG keys**.
    - Click on **New SSH key**.
    - Enter a title (e.g., "My Laptop") and paste your SSH key into the "Key" field.
    - Click **Add SSH key**.

## Step 5: Test SSH Connection to GitHub
Now, test your SSH connection to GitHub with:
```bash
ssh -T git@github.com
```

If everything is set up correctly, you’ll see a message like:
```
Hi <username>! You've successfully authenticated, but GitHub does not provide shell access.
```

## Step 6: Update Git Remote URL to Use SSH
Finally, update your Git repository to use the SSH URL instead of HTTPS:
```bash
git remote set-url origin git@github.com:HamidrezaDadafarid/Golang.git
```
Now, you can push and pull without entering your username and password each time:
```bash
git push -u origin master
```
